package tracer

import (
    "bytes"
    "fmt"

    bpf "github.com/pizchen/gobpf/elf"
)

type PerfOutputMap interface {
    Name() string
    MaxEntries() int
    PageCount() int
    Callback([]byte)
}

type Tracer struct {
    *bpf.Module
    peMap    []*bpf.PerfMap
    stopChan []chan struct{}
}

func NewTracer(pom []PerfOutputMap, bpfProg []byte) (*Tracer, error) {

    reader := bytes.NewReader(bpfProg)

    m := bpf.NewModuleFromReader(reader)
    if m == nil {
        return nil, fmt.Errorf("BPF not supported")
    }

    sectionParams := make(map[string]bpf.SectionParams)
    for _, mp := range pom {
        sp := bpf.SectionParams{}
        if mp.MaxEntries() > 0 {
            sp.MapMaxEntries = mp.MaxEntries()
        }
        if mp.PageCount() > 0 {
            sp.PerfRingBufferPageCount = mp.PageCount()
        }
        sectionParams[mp.Name()] = sp
    }

    err := m.Load(sectionParams)
    if err != nil {
        return nil, err
    }

    err = m.EnableKprobes(1)
    if err != nil {
        return nil, err
    }

    for tp := range m.IterTracepointProgram() {
        err = m.EnableTracepoint(tp.Name)
        if err != nil {
            return nil, err
        }
    }

    pm := make([]*bpf.PerfMap, 0)
    stopChan := make([]chan struct{}, 0)

    for _, mp := range pom {

        evChan := make(chan []byte)

        evMap, err := initChan(m, mp.Name(), evChan)
        if err != nil {
            return nil, fmt.Errorf("failed to init perf map: %s", err)
        }

        tChan := make(chan struct{})
        go func() {
            for {
                select {
                case <-tChan:
                    // On stop, stopChan will be closed but the other channels will
                    // also be closed shortly after. The select{} has no priorities,
                    // therefore, the "ok" value must be checked below.
                    return
                case data, ok := <-evChan:
                    if !ok {
                        return // see explanation above
                    }
                    mp.Callback(data)
                }
            }
        }()

        pm = append(pm, evMap)
        stopChan = append(stopChan, tChan)
    }

    return &Tracer{m, pm, stopChan}, nil
}

func (t *Tracer) EnableUprobes(path, name string) error {
    //var err error
    //for up := range t.M.IterUprobes() {
    //    if up.Name == name {
    //    }
    //}
    return nil
}

func (t *Tracer) Start() {
    for _, pm := range t.peMap {
        pm.PollStart()
    }
}

func (t *Tracer) Stop() {
    for i, pm := range t.peMap {
        close(t.stopChan[i])
        pm.PollStop()
    }
    t.Close()
}

func initChan(m *bpf.Module, mapName string, c chan []byte) (*bpf.PerfMap, error) {

    pm, err := bpf.InitPerfMap(m, mapName, c, nil)
    if err != nil {
        return nil, fmt.Errorf("error initializing perf map for %q: %v", mapName, err)
    }

    return pm, nil
}

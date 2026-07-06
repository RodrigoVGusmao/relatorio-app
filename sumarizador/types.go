package main

type JData map[string]any
func (src JData) Clone() JData {
    if src == nil {
        return nil
    }

    cloned := make(JData, len(src))
    for key, val := range src {
        cloned[key] = val
    }
    
    return cloned
}

func toFloat64(val any) (float64, bool) {
	v, ok := val.(float64)
	return v, ok
}

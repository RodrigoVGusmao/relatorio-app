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
	switch v := val.(type) {
	case float64:
		return v, true
	case int:
		return float64(v), true
	case int64:
		return float64(v), true
	case float32:
		return float64(v), true
	case uint32:
		return float64(v), true
	}
	return 0, false
}

package check_x

import (
	"bytes"
	"fmt"
	"strconv"
	"sync"
)

type performanceData map[string]interface{}

var p = []performanceData{}
var pMutex = &sync.Mutex{}

//NewPerformanceData adds a Performancedata object which can be expanded with further information
func NewPerformanceData(label string, value float64) *performanceData {
	return NewPerformanceDataString(label, strconv.FormatFloat(value, 'f', -1, 64))
}

//NewPerformanceDataString adds a Performancedata object which can be expanded with further information
func NewPerformanceDataString(label, value string) *performanceData {
	pMutex.Lock()
	p = append(p, performanceData{"label": label, "value": value})
	newOne := &(p[len(p)-1])
	pMutex.Unlock()
	return newOne
}

//Unit adds an unit string to the performancedata
func (p performanceData) Unit(unit string) performanceData {
	p["unit"] = unit
	return p
}

//Warn adds the threshold to the performancedata
func (p performanceData) Warn(warn *Threshold) performanceData {
	p["warn"] = warn
	return p
}

//Crit adds the threshold to the performancedata
func (p performanceData) Crit(crit *Threshold) performanceData {
	p["crit"] = crit
	return p
}

//Min adds the float64 to the performancedata
func (p performanceData) Min(min float64) performanceData {
	p["min"] = min
	return p
}

//Min adds the float64 to the performancedata
func (p performanceData) Max(max float64) performanceData {
	p["max"] = max
	return p
}

//toString prints this performancedata
func (p performanceData) toString() string {
	var toPrint bytes.Buffer

	toPrint.WriteString(fmt.Sprintf("'%s'=%s", p["label"], p["value"]))
	if unit, ok := p["unit"]; ok {
		toPrint.WriteString(unit.(string))
	}
	toPrint.WriteString(";")
	addThreshold := func(key string) {
		if value, ok := p[key]; ok && value != nil {
			if t := value.(*Threshold); t != nil {
				toPrint.WriteString(t.input)
			}
		}
		toPrint.WriteString(";")
	}
	addThreshold("warn")
	addThreshold("crit")

	addFloat := func(key string) {
		if value, ok := p[key]; ok {
			toPrint.WriteString(strconv.FormatFloat(value.(float64), 'f', -1, 64))
		}
	}
	addFloat("min")
	toPrint.WriteString(";")
	addFloat("max")

	return toPrint.String()
}

//PrintPerformanceData prints all performancedata
func PrintPerformanceData() string {
	var toPrint bytes.Buffer
	pMutex.Lock()
	for _, perfData := range p {
		toPrint.WriteString(perfData.toString())
		toPrint.WriteString(" ")
	}
	pMutex.Unlock()
	return toPrint.String()
}

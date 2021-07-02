package vrchat

func (v *VRCGO) RegisterEvent(name string, callback func(interface{})) {
	list := v.events[name]
	list = append(list, callback)
	v.events[name] = list
}

func (v *VRCGO) CallEvent(name string, param interface{}) {
	list := v.events[name]
	for _, callback := range list {
		callback(param)
	}
}

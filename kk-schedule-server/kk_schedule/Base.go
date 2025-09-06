package kk_schedule

func (x *PBRegisterService) Check() error {
	if x.Target == "" {
		return ErrTargetEmpty
	}
	if x.ServiceName == "" {
		return ErrServiceNameEmpty
	}
	return nil
}

func (x *PBRegisterJob) Check() error {
	if x.FuncName == "" {
		return ErrFuncNameEmpty
	}
	return nil
}

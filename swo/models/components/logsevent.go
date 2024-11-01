// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type LogsEvent struct {
	ID       string `json:"id"`
	Time     string `json:"time"`
	Message  string `json:"message"`
	Hostname string `json:"hostname"`
	Severity string `json:"severity"`
	Program  string `json:"program"`
}

func (o *LogsEvent) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LogsEvent) GetTime() string {
	if o == nil {
		return ""
	}
	return o.Time
}

func (o *LogsEvent) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *LogsEvent) GetHostname() string {
	if o == nil {
		return ""
	}
	return o.Hostname
}

func (o *LogsEvent) GetSeverity() string {
	if o == nil {
		return ""
	}
	return o.Severity
}

func (o *LogsEvent) GetProgram() string {
	if o == nil {
		return ""
	}
	return o.Program
}

package playrlo

func Some[T any](value T) *T {
	return &value
}

type CompileRequest struct {
	Target           CompileTarget     `json:"target"`
	AssemblyFlavor   *AssemblyFlavor   `json:"assemblyFlavor"`
	DemangleAssembly *DemangleAssembly `json:"demangleAssembly"`
	ProcessAssembly  *ProcessAssembly  `json:"processAssembly"`
	Channel          Channel           `json:"channel"`
	Mode             Mode              `json:"mode"`
	Edition          *Edition          `json:"edition"`
	CrateType        CrateType         `json:"crateType"`
	Tests            bool              `json:"tests"`
	Backtrace        bool              `json:"backtrace"`
	Code             string            `json:"code"`
}

type FormatRequest struct {
	Code    string   `json:"code"`
	Edition *Edition `json:"edition"`
}

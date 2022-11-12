package playrlo

type AssemblyFlavor string

const (
	Att   AssemblyFlavor = "att"
	Intel AssemblyFlavor = "intel"
)

type DemangleAssembly string

const (
	Demangle DemangleAssembly = "demangle"
	Mangle   DemangleAssembly = "mangle"
)

type ProcessAssembly string

const (
	Filter ProcessAssembly = "filter"
	Raw    ProcessAssembly = "raw"
)

type CompileTarget string

const (
	ASM    CompileTarget = "asm"
	LLVMIR CompileTarget = "llvm-ir"
	MIR    CompileTarget = "mir"
	HIR    CompileTarget = "hir"
	WASM   CompileTarget = "wasm"
)

type Channel string

const (
	Stable  Channel = "stable"
	Beta    Channel = "beta"
	Nightly Channel = "nightly"
)

type CrateType string

const (
	Binary    CrateType = "bin"
	Lib       CrateType = "lib"
	Dylib     CrateType = "dylib"
	Rlib      CrateType = "rlib"
	Staticlib CrateType = "staticlib"
	Cdylib    CrateType = "cdylib"
	ProcMacro CrateType = "proc-macro"
)

type Mode string

const (
	Debug   Mode = "debug"
	Release Mode = "release"
)

type Edition string

const (
	Rust2015 Edition = "2015"
	Rust2018 Edition = "2018"
	Rust2021 Edition = "2021"
)

package lint

func RunBoltLinters(f *file) {
	logImport(f)
}

// Examines the imports to ensure no file imports "log" unless it is imported as "std_log"
func logImport(f *file) {
	for i, is := range f.f.Imports {
		_ = i
		if is.Path != nil && is.Path.Value == "\"log\"" && !f.isTest() &&
			(is.Name == nil || is.Name.Name != "std_log") {

			f.errorf(is, 1, link(""), category("imports"),
				"Do not import log directly, use common/log instead. If necessary name import std_log")
		}
	}
}

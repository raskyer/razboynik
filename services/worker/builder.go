package worker

import "github.com/eatbytes/razboy/core"

func BuildShellConfig(method, scope string) *core.SHELLCONFIG {
	return &core.SHELLCONFIG{
		Method: method,
		Scope:  scope,
		Cmd:    "",
	}
}

func BuildPHPConfig(raw, upload bool) *core.PHPCONFIG {
	return &core.PHPCONFIG{
		Cmd:    "",
		Raw:    raw,
		Upload: upload,
	}
}

func BuildServerConfig(url, method, parameter, key string, raw bool) *core.SERVERCONFIG {
	return &core.SERVERCONFIG{
		Url:       url,
		Method:    method,
		Parameter: parameter,
		Key:       key,
		Raw:       raw,
	}
}

func BuildRequest(shl *core.SHELLCONFIG, php *core.PHPCONFIG, srv *core.SERVERCONFIG) *core.REQUEST {
	return &core.REQUEST{
		SHLc: shl,
		PHPc: php,
		SRVc: srv,
	}
}

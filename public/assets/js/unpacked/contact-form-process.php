<?php
    $to = "christinabeymer@gmail.com"; // put your email address here
    $from = $_REQUEST['email'];
    $name = $_REQUEST['name'];
    $headers = "From: $from";
    $subject = "$contactsubject from: $name @ $from"; // subject and the name and email of the person

    $fields = array();
    $fields{"name"} = "Name";
    $fields{"email"} = "Email";
    $fields{"contactsubject"} = "Subject";
    $fields{"phone"} = "Phone";
    $fields{"message"} = "Message";

    $body = "Here is what was sent:\n\n"; foreach($fields as $a => $b){   $body .= sprintf("%20s: %s\n\n",$b,$_REQUEST[$a]); }

    $send = mail($to, $subject, $body, $headers);

?>
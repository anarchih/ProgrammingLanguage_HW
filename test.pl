open(my $str, '<:encoding(UTF-8)', "test.html");


$row = <$str>;
$i = 0;
while (my $row = <$str>){
    # if ($i > 20){
    #    last;
    # }
    $i += 1;
    if($row =~ /\s{10,100}<a\shref=".+>.*<\/a>\s+/){
        print $row;
    }
}

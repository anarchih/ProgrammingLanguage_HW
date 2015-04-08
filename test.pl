use LWP::Simple;
use Encode;
use utf8;
use open ':std', ':encoding(UTF-8)';
$url = "http://www.atmovies.com.tw/showtime/theater_t06608_a06.html";
getstore($url, 'download.html');
open(my $str, '<:encoding(UTF-8)', "download.html");
$i = 0;
while (my $row = <$str>){
    $i += 1;
    # for name
    if($row =~ /(\s{10,100}<a\shref=".+>.*<\/a>\s+)/){
        print $row =~ /.*>(.*)<.*/;
        print "\n";
        
    }
    # print $row;
    # for time
    if($row =~ /\s+(<LI>.+<\/LI>)+/){
        my @test = $row =~ /(\d+：\d+)/g;
        foreach (@test) {
              print "$_\n";
        }
    }
}

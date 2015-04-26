use LWP::Simple;
use utf8;
use open ':std', ':encoding(UTF-8)';
@urls = ("http://www.atmovies.com.tw/showtime/theater_t06608_a06.html",
         "http://www.atmovies.com.tw/showtime/theater_t06607_a06.html",
         "http://www.atmovies.com.tw/showtime/theater_t06609_a06.html");
for my $url (@urls){
    getstore($url, 'download.html');
    open(my $str, '<:encoding(UTF-8)', "download.html");
    while (my $row = <$str>){
        # for name
        if($row =~ /(\s{10,100}<a\shref=".+>(.*)<\/a>\s+)/){
            print "$2\n";
        }
        if($row =~ /\s+(<LI>.+)+/){
            my @test = $row =~ /(\d+：\d+)/g;
            foreach (@test) {
                print "$_\n";
            }
        }
    }
}

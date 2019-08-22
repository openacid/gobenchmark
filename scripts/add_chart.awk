BEGIN {
    sym = "------------------------------------------------------------"
}
{
    ns = $(NF-1)
    ns *= 5
    ns = int(ns)
    # remove bench times
    gsub("	\d[0	]*", "", $0)
    print $0" " substr(sym, 0, ns)
}

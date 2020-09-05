d = Dict('a'=>1, 'b'=>2)
println(d)

println(get(d,"a", -1))

println(merge(d, Dict('c'=>99)))
println(d)

merge!(d, Dict('c'=>99))
println(d)


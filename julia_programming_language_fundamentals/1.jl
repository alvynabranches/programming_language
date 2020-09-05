# Create a structure
struct Person
    name::AbstractString
    id::Int
end

people = Person[]
push!(people, Person("Alvyn Abranches", 10001))
push!(people, Person("Anil Gadwal", 10002))
println(people)
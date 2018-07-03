Counter = {}

function Counter:new( v )
    c = {value=v or 0}
    setmetatable(c, self)
    self.__index = self 
    return c
end

function Counter:get() 
    return self.value
end 

counter = Counter:new(100)
for i=1,10 do 
    print(counter:incr(i))
end 

print(counter:get())


cases, _ = io.read("*number", "*line")
while cases > 0 do
    cases = cases - 1
    name = io.read("*line")
    print("Hello, " .. name .."!")
end

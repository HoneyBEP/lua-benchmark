function benchmark(depth)
    print('print `depth` from Lua: ' .. depth)
    return  depth + 10
end

function countFromGo()
    print(concat("test", "1234"))
    print('echo `count` in Lua')
    print(count())
end

function showDate()
    print(os.date("%a %b %d %H:%M:%S"))
end

function showVariable()
    local goVariable = goGetVariable()
    print(goVariable)
    goVariable = goVariable .. " is a variable"
    return goVariable
end

function runPeopleDemo()
    local p = person.new("Steeve")
    print(p:name()) -- "Steeve"
    p:name("Alice")
    print(p:name()) -- "Alice"
end

--function getGoVariable()
--    return 'a variable'
--end
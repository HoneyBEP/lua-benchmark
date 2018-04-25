function benchmark(depth)
    print('print `depth` from Lua: ' .. depth)
    return  depth + 10
end

function countFromGo()
    print('echo `count` in Lua')
    count()
end
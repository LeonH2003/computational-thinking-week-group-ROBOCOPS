-- Read data from `data4.csv`
local lines = {}
for line in io.lines("challenge_day4/Julia_step3/data4.csv") do
    table.insert(lines, line)
end

-- Skip header
table.remove(lines, 1)

-- Prepare output table
local people = {}

-- Process each line
for _, line in ipairs(lines) do
    local name, tech, soft, bus, creative, academic = line:match("([^,]+),([^,]+),([^,]+),([^,]+),([^,]+),([^,]+)")

    -- Count occurrences of each skill level
    local counts = {super=0, good=0, middle=0, low=0}
    for _, skill in ipairs({tech, soft, bus, creative, academic}) do
        counts[skill] = (counts[skill] or 0) + 1
    end

    -- Determine final summary
    local finalSummary
    if counts.super > 0 then
        finalSummary = "super"
    elseif counts.good >= 2 then
        finalSummary = "good"
    elseif counts.middle >= 3 then
        finalSummary = "middle"
    else
        finalSummary = "low"
    end

    table.insert(people, {name, tech, soft, bus, creative, academic, finalSummary})
end

-- Write output to `data5.txt`
local out = io.open("challenge_day4/Lua_step4/data5.txt", "w")
out:write("Name,Technical Skills,Soft Skills,Business Skills,Creative Skills,Academic Skills,Summary\n")
for _, entry in ipairs(people) do
    out:write(table.concat(entry, ',') .. "\n")
end
out:close()

print("Done! Generated challenge_day4/Lua_step4/data5.txt")

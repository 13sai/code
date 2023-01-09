local data="{\"id\":1}"
wrk.headers["Content-Type"] =  "application/jason"
wrk.method  =  "POST"
function  request()
    return wrk.format('POST', nil, nil, data)
end
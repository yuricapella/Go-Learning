aula 03
Value - type - bytes
1234567890 - int64 - 00000000499602d2
3.141592653589793 - float64 - 400921fb54442d18
"<<Serialization>>" - string - c2ab53657269616c697a6174696f6ec2bb

como organizar os bytes
left or right side
porque se chama big endian
guliver travels
little endian, big endian, o povo de lá da historia?
esta andando no pequeno ou no grande endian?
string

descricao do video:
Why do we need serialization? 

Serialization is the process of taking a structure or data type from the language you're working in and converting it to a series of bytes (Marshalling). When we want to transmit integers from one place to another we must serialize our data to be read by the machine. Let's take a closer look.


---
aula 04
Practical Serialization In Go: HTTP Handler with JSON data

const (
    jsonCtype = "aplication/json"
)

func addHandler(w http.ResponseWriter, r *http.Request){
    //Step 1: De-serialize (unmarshal)

    var rec weather.Record
    defer r.Body.Close()

    if err := json.NewDecoder(r.Body).Decode(&rec); err!= nil {
        log.Printf("unmarshal: %s", err)
        http.Error(w, err.Error(), httpStatusBadRequest)
        return
    }

    //Step 2: Work
    log.Printf("adding %#v", rec)
    n:= weather.AddRecord(rec)


    // Step 3: Serialize (marshal)
    resp := map[string]interface{}{
        "ok": true,
        "num_records": n,
    }
    w.Header().Set("Contend-Type", jsonCtype)
    if r:= json.NewEncoder(w).Encode(resp); err != nil {
        // cant notify client on error once you start writing to w
        log.Printf("marshal: %s", err)
    }
}


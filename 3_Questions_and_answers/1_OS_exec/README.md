# OS Exec

## How to add `cat` in PATH - on Windows

1. Install Git Bash
2. Assuming we are using default directory to install add `C:\Program Files\Git\usr\bin` to `PATH`

## How to test `GetData` with own data

```
func GetData() string {
	cmd := exec.Command("cat", "msg.xml")

	out, _ := cmd.StdoutPipe()
	var payload Payload
	decoder := xml.NewDecoder(out)

	// these 3 can return errors but I'm ignoring for brevity
	cmd.Start()
	decoder.Decode(&payload)
	cmd.Wait()

	return strings.ToUpper(payload.Message)
}
```

### Option 1

Add one more argument in `GetData(file path)` and prepare a list of test files for comparison.

### Option 2

The alternative is to split logic into:

```
func getXMLFromCommand() io.Reader {
	cmd := exec.Command("cat", "msg.xml")

	out, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalln(err)
	}
	data, err := ioutil.ReadAll(out)
	if err != nil {
		log.Fatalln(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatalln(err)
	}

	return bytes.NewReader(data)
}

```

and

```
func GetData(data io.Reader) string {
	var payload Payload
	xml.NewDecoder(data).Decode(&payload)
	return strings.ToUpper(payload.Message)
}
```

we can use `strings.NewReader` to create our own data for testing

```
func TestGetData(t *testing.T) {
	input := strings.NewReader(`
<payload>
	<message>Cats are the best animal</message>
</payload>`)
	got := GetData(input)
	want := "CATS ARE THE BEST ANIMAL"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

```
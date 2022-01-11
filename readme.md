# In memory of my dog Safira.

![logo of Safira, my dog](./logo.png)

The tool `safira` replaces `pattern_offset.rb` and `pattern_create.rb` from the `MetaSploit Framework`.

# Compile
You can run one of the `compile` scripts accodringly to your OS target.

---

Linux:

```bash
bash ./compile.sh
```

Windows: 
```batch
.\compile.bat
```
(or double click on the script).

---

In case you're not neither of one of the two, you could run the following command:

```bash
go build -ldflags "-s -w" -o "./bin/safira" "./src/safira.go"
```

# Usage

To create a pattern, you might supply the lenght of it:

```bash
safira -lenght 40
# Aa0Aa1Aa2Aa3Aa4Aa5Aa6Aa7Aa8Aa9Ab0Ab1Ab2
```

After retreiving the data that you see on the `EIP`, pass it to the `-find` argument.

```bash
safira -lenght 40 -find 41306241
# Matching at 30 bytes
```

# Why
- I dislike Ruby
- I don't want to install Metasploit Framework on Windows
- I like making programs in Go
- My dog died because she ate some poisonous stuff that some uncivilized human dropped on the open field on purpose. 

# Your code/program sucks, what can I do about it?
- Make a push request.
- Fork it and edit as you like / make your own version of it.
- Absolutely nothing.

If none of the options are appealing you, you can contact me at the following number: `605 475-6968`.



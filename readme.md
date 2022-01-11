# In memory of my dog Safira.

![logo of Safira, my dog](./logo.png)

The tool `safira` replaces `pattern_offset.rb` and `pattern_create.rb` from the `MetaSploit Framework`.

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

# Your code/program sucks, what can I do about it?
- Make a push request.
- Nothing

If none of the options are appealing you, you can contact me at the following number: `605 475-6968`.

Commit messages in `thetaRadixDebuffer` follow a specific set of conventions, which we discuss in this section.

Here is an example of a good one:
```
math: improve Sin, Cos and Tan precision for very large arguments

The existing implementation has poor numerical properties for
large arguments, so use the McGillicutty algorithm to improve
accuracy above 1e10.

The algorithm is described at https://wikipedia.org/wiki/McGillicutty_Algorithm

Fixes #159
```
## First line
The first line of the change description is conventionally a short one-line summary of the change, prefixed by the primary affected package.

A rule of thumb is that it should be written so to complete the sentence "This change modifies `thetaRadixDebuffer` to _____." That means it does not start with a capital letter, is not a complete sentence, and actually summarizes the result of the change.

Follow the first line by a blank line.

## Main content
The rest of the description elaborates and should provide context for the change and explain what it does. Write in complete sentences with correct punctuation, just like for your comments in Go. Don't use HTML, Markdown, or any other markup language.

Add any relevant information, such as benchmark data if the change affects performance.

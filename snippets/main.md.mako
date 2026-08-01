<%!
    import pydmt.helpers.git
    file_number = pydmt.helpers.git.count_files("src/**/*.go")
%>${"##"} Number of examples

Currently there are ${file_number} examples in this repo.

# tree-sitter-dream

Tree-sitter grammar for the Dream language - a Rust-like syntax with Erlang/BEAM concurrency semantics.

## Features

- Rust-like syntax (functions, structs, enums, traits, impl blocks)
- Erlang-style concurrency primitives (`spawn`, `send !`, `receive`)
- Pattern matching with guards
- Atoms (`:ok`, `:error`)
- Pipe operator (`|>`)
- Bitstring expressions (`<<value:size/specifiers>>`)

## Building

```bash
npm install
npx tree-sitter generate
```

## Testing the Parser

```bash
npx tree-sitter parse path/to/file.dream
```

## Neovim Setup

### 1. Install nvim-treesitter

Make sure you have [nvim-treesitter](https://github.com/nvim-treesitter/nvim-treesitter) installed.

### 2. Register the parser

Add this to your Neovim config:

```lua
local parser_config = require("nvim-treesitter.parsers").get_parser_configs()

parser_config.dream = {
  install_info = {
    url = "https://github.com/scrogson/tree-sitter-dream",
    files = { "src/parser.c" },
    generate_requires_npm = true,
  },
  filetype = "dream",
}
```

### 3. Set up filetype detection

Create `~/.config/nvim/ftdetect/dream.vim`:

```vim
au BufRead,BufNewFile *.dream set filetype=dream
```

Or in Lua:

```lua
vim.filetype.add({
  extension = {
    dream = "dream",
  },
})
```

### 4. Copy query files

```bash
mkdir -p ~/.config/nvim/queries/dream
cp queries/*.scm ~/.config/nvim/queries/dream/
```

### 5. Install the parser

In Neovim:

```vim
:TSInstall dream
```

### 6. Verify

Open a `.dream` file and run:

```vim
:InspectTree
```

## Example Syntax

```rust
mod counter {
    pub fn start() -> Pid {
        spawn || {
            loop(0)
        }
    }

    fn loop(count: int) {
        receive {
            :increment => loop(count + 1),
            (:get, sender) => {
                sender ! count;
                loop(count)
            }
        }
    }
}
```

## Syntax Highlighting

The grammar provides highlighting for:

- Keywords (`fn`, `let`, `if`, `match`, `struct`, `enum`, `spawn`, `receive`, etc.)
- Operators (`+`, `-`, `*`, `/`, `==`, `!=`, `&&`, `||`, `!`, `|>`, etc.)
- Bitstring syntax (`<<`, `>>`, `big`, `little`, `signed`, `unsigned`, etc.)
- Literals (integers, strings, atoms, booleans)
- Types (type identifiers, primitives)
- Functions (definitions and calls)
- Comments (line `//` and block `/* */`)

## License

MIT

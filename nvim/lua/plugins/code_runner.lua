return {
  {
    "CRAG666/code_runner.nvim",
    config = function()
      require("code_runner").setup({
        filetype = {
          haskell = {
            "cd $dir && ",
            "ghc $fileName &&",
            "$dir/$fileNameWithoutExt",
          },
          go = {
            "cd $dir &&",
            "go run $fileName",
          },
          java = {
            "cd $dir &&",
            "java $fileName &&",
            "java $fileNameWithoutExt",
          },
          kotlin = {
            "cd $dir &&",
            "kotlinc $fileName -include-runtime -d  $fileName.jar &&",
            "java -jar $fileName.jar",
          },
          dart = "dart run ",
          python = "python3 -u",
          typescript = "bun run",
          rust = {
            "cd $dir &&",
            "rustc $fileName &&",
            "$dir/$fileNameWithoutExt",
          },
          lua = "lua",
          zig = {
            "cd $dir &&",
            "zig run $fileName",
          },
        },
      })
    end,

    keys = {
      {

        "<leader>r",
        ":RunCode<CR>",
      },
      {
        "<leader>R",
        ":RunClose<CR>",
      },
    },
  },
}

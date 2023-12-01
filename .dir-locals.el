((nil . ((format-all-formatters . (("Go" gofmt goimports)))))
 (go-ts-mode
  (eval add-hook 'before-save-hook #'format-all-buffer nil 'local)))

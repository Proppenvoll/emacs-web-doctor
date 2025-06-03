(defun extract-answers (input-string)
  "Extracts answers from a string that uses \n\n as a separator."
  (seq-filter
   (lambda (element) element)
   (seq-map-indexed
    (lambda (element index)
      (if (= (mod index 2) 0)
          element
        nil))
    (split-string input-string "\n\n"))))

(defun read-stdin ()
  (with-temp-buffer
    (insert-file-contents-literally "/dev/stdin")
    (buffer-string)))

(progn
  (doctor)
  (let ((answers (extract-answers (read-stdin))))
    (dolist (answer answers)
      (insert answer "\n")
      (doctor-read-print))
    t)
  (princ (buffer-string)))

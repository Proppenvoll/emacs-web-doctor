(ert-deftest extract-answers-test ()
  (should (equal (extract-answers "") '("")))
  (should (equal (extract-answers "\n\n") '("")))
  (should (equal (extract-answers "\n\nquestion") '("")))
  (should (equal (extract-answers "\n\nquestion\n\n") '("" "")))
  (should (equal (extract-answers "answer") '("answer")))
  (should
   (equal
    (extract-answers
     "answer\n\nquestion\n\nanother-answer")
    '("answer" "another-answer")))
  (should
   (equal
    (extract-answers
     "line1\nline2\n\n\n\nline3\nline4")
    '("line1\nline2" "line3\nline4"))))

(defun split (sequence)
  (let ((mid (ceiling (length sequence) 2)))
    (values (subseq sequence 0 mid)
            (subseq sequence mid nil)
    )
  )
)

(defun merge-two-list (left right)
  (cond 
    ((= 0 (length left))
      right
    )
    ((= 0 (length right))
      left
    )
    (t
      (if (< (car left) (car right))
        (cons (car left) (merge-two-list (cdr left) right))
        (cons (car right) (merge-two-list left (cdr right)))
      )
    )
  )
)

(defun mergeSort (l)
  (cond ((= 1 (length l))
         l)
        (t
         (multiple-value-bind (x y) (split l)
         (merge-two-list (mergeSort x) (mergeSort y)))
         )
  )
)

(read)
(
  (lambda (a) 
    (loop for x in a do (princ x)(princ #\ )))
  (mergeSort 
    (with-input-from-string 
      (s (read-line)) 
      (loop for x = (read s nil :end) until (eq x :end) collect x)
    )
  )
)

(defun split (sequence)
  (let ((mid (ceiling (length sequence) 2)))
    (values (subseq sequence 0 mid)
            (subseq sequence mid nil)
    )
  )
)

(defun merge-two-list (left right)
  (let ((r '()))
    (loop while (and left right) do
      (if (< (car left) (car right))
        (progn
          (setq r (append r (list (car left))))
          (setq left (cdr left)))
        (progn
          (setq r (append r (list (car right))))
          (setq right (cdr right)))
      )
    )
    (if left (setq r (append r left)))
    (if right (setq r (append r right)))
    r
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

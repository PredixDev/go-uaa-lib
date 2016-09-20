!(\
       /should have comment or be unexported.*\(golint\)/ \
    || /should have comment \(or a comment on this block\) or be unexported.*\(golint\)/ \
    || /comment on exported type .+ should be of the form.*\(golint\)/ \
    || /error return value not checked \(defer .*\) \(errcheck\)/ \
    || /_test\.go.* error return value not checked \(w.Write.*\) \(errcheck\)/ \
    || /libfakes\/fake_[A-Za-z0-9_]+.go.*/ \
)

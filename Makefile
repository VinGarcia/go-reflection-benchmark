
TIME=1s

all: reading-tags for-each

reading-tags:
	go test -bench=. ./readingtags/... -benchtime=$(TIME)

for-each:
	go test -bench=. ./genericforeach/... -benchtime=$(TIME)

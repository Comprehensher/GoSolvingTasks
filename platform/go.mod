module platform

go 1.22.5

replace platform/logging => ./logging

replace platform/config => ./config

require (
	platform/config v0.0.0-00010101000000-000000000000
	platform/logging v0.0.0-00010101000000-000000000000
)

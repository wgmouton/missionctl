package main

import "fmt"

type Label interface {
	SetTag(tag string)
	SetRevision(revision uint8)

	getLabel() string

	IncrementRevision()
}

type ReleaseCandidate struct {
	Tag      string
	Revision uint8
}

func (rc *ReleaseCandidate) SetTag(tag string) {
	rc.Tag = tag
}

func (rc *ReleaseCandidate) SetRevision(revision uint8) {
	rc.Revision = revision
}

func (rc *ReleaseCandidate) getLabel() string {
	return fmt.Sprintf("%s.%v", rc.Tag, rc.Revision)
}

func (rc *ReleaseCandidate) IncrementRevision() {
	rc.Revision++
}

type Version struct {
	Major  uint8
	Minor  uint8
	BugFix uint8
	Label  Label
	Meta   string
}

func (v Version) GetVersion() string {
	return fmt.Sprintf("%v.%v.%v-%s+%s", v.Major, v.Minor, v.BugFix, v.Label.getLabel(), v.Meta)
}

func ParseVersion(version string) Version {

	var label Label = &ReleaseCandidate{
		Tag:      "rc",
		Revision: 1,
	}

	return Version{
		Major:  1,
		Minor:  2,
		BugFix: 3,
		Label:  label,
		Meta:   "build.1",
	}
}

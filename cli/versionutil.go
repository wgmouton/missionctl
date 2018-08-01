package main

import "fmt"
import "regexp"
import "strconv"

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
	Major uint8
	Minor uint8
	Patch uint8
	Label Label
	Meta  string
}

func (v Version) GetVersion() string {
	var label string
	if v.Label == nil {
		label = ""
	} else {
		label = "-" + v.Label.getLabel()
	}

	var meta string
	if v.Meta == "" {
		meta = ""
	} else {
		meta = "+" + v.Meta
	}

	return fmt.Sprintf("%v.%v.%v%s%s", v.Major, v.Minor, v.Patch, label, meta)
}

func (v *Version) IncrementMajor() {
	v.Major++
	v.Minor = 0
	v.Patch = 0
	v.Label.SetRevision(0)
}

func (v *Version) IncrementMinor() {
	v.Minor++
	v.Patch = 0
	v.Label.SetRevision(0)
}

func (v *Version) IncrementPatch() {
	v.Patch++
	v.Label.SetRevision(0)
}

func (v *Version) IncrementLabelRevision() {
	v.Label.IncrementRevision()
}

func (v *Version) SetLabel(label Label) {
	v.Label = label
}

func (v *Version) SetLabelFromString(label string) {
	v.Label = parseLabel(label)
}

func (v *Version) RemoveLabel() {
	v.Label = nil
}

func (v *Version) SetMetaData(meta string) {
	v.Meta = meta
}

func (v *Version) RemoveMeta() {
	v.Meta = ""
}

func parseVersionDigit(versionDigit string) uint8 {
	digit, err := strconv.ParseInt(versionDigit, 8, 8)
	if err != nil {
		panic(fmt.Sprintf("parse error: %s is not an int: %s", versionDigit, err.Error()))
	}

	return uint8(digit)
}

func parseLabel(label string) Label {
	releaseCandidateRegex, _ := regexp.Compile("rc\\.(\\d+)")

	switch {
	case releaseCandidateRegex.MatchString(label):
		matches := releaseCandidateRegex.FindAllStringSubmatch(label, -1)
		return &ReleaseCandidate{
			Tag:      "rc",
			Revision: parseVersionDigit(matches[0][1]),
		}
	default:
		return nil
	}
}

func ParseVersion(version string) Version {

	regex, _ := regexp.Compile("(\\d+)\\.(\\d+)\\.(\\d+)(?:-(\\w+(?:\\.\\w+)*))?(?:\\+(\\w+(?:\\.\\w+)*))?")
	matches := regex.FindAllStringSubmatch(version, -1)

	major, minor, patch := parseVersionDigit(matches[0][1]), parseVersionDigit(matches[0][2]), parseVersionDigit(matches[0][3])
	label, meta := parseLabel(matches[0][4]), matches[0][5]

	return Version{
		Major: major,
		Minor: minor,
		Patch: patch,
		Label: label,
		Meta:  meta,
	}
}

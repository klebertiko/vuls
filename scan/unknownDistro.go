/* Vuls - Vulnerability Scanner
Copyright (C) 2016  Future Architect, Inc. Japan.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package scan

// inherit OsTypeInterface
type unknown struct {
	base
}

func (o *unknown) checkIfSudoNoPasswd() error {
	return nil
}

func (o unknown) checkDependencies() error {
	return nil
}

func (o *unknown) install() error {
	return nil
}

func (o *unknown) checkRequiredPackagesInstalled() error {
	return nil
}

func (o *unknown) scanPackages() error {
	return nil
}

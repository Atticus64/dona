%global goipath         github.com/atticus64/dona
%global gomodulesmode   GO111MODULE=on
Version:                0.0.6

%gometa

%global golicenses      LICENSE
%global godocs          README.md

Name:           dona
Release:        1%{?dist}
Summary:        CLI to manage your dot files
License:        Apache-2.0 AND BSD-3-Clause AND MIT
URL:            %{gourl}
Source0:        https://github.com/Atticus64/dona/archive/refs/tags/v%{version}.tar.gz
Source1:        dona-%{version}-vendor.tar.gz

BuildRequires:  go-rpm-macros
BuildRequires:  golang >= 1.24

ExclusiveArch:  %{golang_arches}

%global debug_package %{nil}

%description
CLI to search across GitHub dot files configurations and save them.

%prep
%goprep -A
tar xf %{SOURCE1} --strip-components=1 -C . --wildcards '*/vendor'

%build
%gobuild -o dona %{goipath}

%install
install -Dpm 0755 dona %{buildroot}%{_bindir}/dona

%files
%license LICENSE
%doc README.md
%{_bindir}/%{name}

%changelog
* Sat Apr 04 2026 Jonathan <jonathanelian64@gmail.com> - 0.0.6-1
- Initial package

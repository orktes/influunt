from setuptools import setup, find_packages, Distribution
import os
import platform
import sys

VERSION_FILE = os.path.join(os.path.dirname(__file__), 'influunt/VERSION')
README_MD_FILE = os.path.join(os.path.dirname(__file__), 'README.md')

class BinaryDistribution(Distribution):
    """Auxilliary class necessary to inform setuptools that this is a
    non-generic, platform-specific package."""
    def has_ext_modules(self):
        return True

try:
    from wheel.bdist_wheel import bdist_wheel as _bdist_wheel
    class bdist_wheel(_bdist_wheel):
        def finalize_options(self):
            _bdist_wheel.finalize_options(self)
            self.root_is_pure = False
except ImportError:
    bdist_wheel = None

LIB_PATH = []

if sys.platform == 'darwin':
    LIB_PATH.append('../build/sharedlib/darwin/amd64/influunt_core.so')
elif sys.platform.startswith('linux'):
    LIB_PATH.append('../build/sharedlib/linux/amd64/influunt_core.so')

setup(name='influunt',
      version=open(VERSION_FILE).read().strip(),
      url='https://github.com/orktes/influunt',
      license='MIT',
      author='Jaakko Lukkari',
      author_email='jaakko.lukkari@gmail.com',
      description='Dataflow programming for Golang and Python',
      packages=find_packages(exclude=['tests']),
      long_description=open(README_MD_FILE).read(),
      zip_safe=False,
      include_package_data=True,
      data_files=[('influunt/lib', LIB_PATH)],
      cmdclass={'bdist_wheel': bdist_wheel},
      distclass=BinaryDistribution,
      )
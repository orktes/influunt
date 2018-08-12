from setuptools import setup, find_packages
import os
import platform
import sys

VERSION_FILE = os.path.join(os.path.dirname(__file__), 'influunt/VERSION')
README_MD_FILE = os.path.join(os.path.dirname(__file__), '../README.md')


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
      )
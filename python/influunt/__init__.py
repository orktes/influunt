# coding: utf-8
"""Influunt: Data flow programming for Python and Go.
"""

import os
import sys
from os.path import dirname, join
sys.path.append(join(dirname(__file__), "lib"))

from .core import *
from .executor import *



VERSION_FILE = os.path.join(os.path.dirname(__file__), 'VERSION')
with open(VERSION_FILE) as f:
    __version__ = f.read().strip()

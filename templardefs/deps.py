'''
dependencies for this project
'''

def populate(d):
    d.packs=[
        'golang-go',
        'golang-doc',
    ]

def getdeps():
    return [
        __file__, # myself
    ]

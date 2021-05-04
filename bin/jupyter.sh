#!/usr/bin bash
conda activate lyt
jupyter lab --allow-root --ip=0.0.0.0&
conda deactivate

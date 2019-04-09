<?php

/**
 * Implements hook_panels_layouts().
 */
// Plugin definition
$plugin = array(
  'title' => t('Dossier details'),
  'category' => t('Charte2016'),
  'icon' => 'dossier.png',
  'theme' => 'dossier',
  'admin css' => 'dossier.css',
  'regions' => array(
    'block1' => t('Block 1'),
    'block2' => t('Block 2'),
    'block3' => t('Block 3'),
    'block4' => t('Block 4'),
    'block5' => t('Block 5'),
    'block6' => t('Block 6'),
  ),
);
